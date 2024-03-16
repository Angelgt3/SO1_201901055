#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>

#define PROC_FILENAME "cpu_so1_1s2024"

static int cpu_proc_show(struct seq_file *m, void *v) {
    char buffer[256];
    struct file *fp;
    ssize_t n;

    fp = filp_open("/proc/stat", O_RDONLY, 0);

    if (fp != NULL) {
        n = kernel_read(fp, buffer, sizeof(buffer) - 1, &fp->f_pos);
        buffer[n] = '\0';

        seq_printf(m, "CPU Info:\n%s\n", buffer);
        filp_close(fp, NULL);
    } else {
        seq_printf(m, "Error: Unable to open /proc/stat\n");
    }

    struct task_struct *task;
    struct task_struct *child;
    for_each_process(task) {
        seq_printf(m, "PID: %d, Name: %s\n", task->pid, task->comm);
        
        if (!list_empty(&task->children)) {
            seq_printf(m, "Children:\n");
            list_for_each_entry(child, &task->children, sibling) {
                seq_printf(m, "  PID: %d, Name: %s\n", child->pid, child->comm);
            }
        }
    }
    
    return 0;
}

static int cpu_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, cpu_proc_show, NULL);
}

static const struct proc_ops cpu_proc_fops = {
    .proc_open    = cpu_proc_open,
    .proc_read    = seq_read,
    .proc_lseek   = seq_lseek,
    .proc_release = single_release,
};

static int __init cpu_module_init(void) {
    proc_create(PROC_FILENAME, 0, NULL, &cpu_proc_fops);
    printk(KERN_INFO "cpu_so1_1s2024 module initialized\n");
    return 0;
}

static void __exit cpu_exit(void) {
    remove_proc_entry(PROC_FILENAME, NULL);
    printk(KERN_INFO "Exiting cpu_so1_1s2024 module\n");
}

module_init(cpu_module_init);
module_exit(cpu_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel");
MODULE_DESCRIPTION("Module CPU");
