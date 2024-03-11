#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/fs.h>
#include <linux/sched.h> 
#include <linux/sched/signal.h> 

#define PROC_FILENAME "cpu_so1_1s2024"

static int cpu_proc_show(struct seq_file *m, void *v) {
    struct task_struct *task;
    int count = 0;
    // Iterar sobre todos los procesos para contarlos
    for_each_process(task) {
        count++;
    }
    // Imprimir el total de procesos en el archivo /proc
    seq_printf(m, "Total number of processes: %d\n", count);
    return 0;
}

static int cpu_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, cpu_proc_show, NULL);
}

static const struct proc_ops cpu_proc_fops = {
    .proc_open   = cpu_proc_open,
    .proc_read   = seq_read,
    .proc_lseek  = seq_lseek,
    .proc_release= single_release,
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
MODULE_DESCRIPTION("Kernel module to obtain CPU and process information");
