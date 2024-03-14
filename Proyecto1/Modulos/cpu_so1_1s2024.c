#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/fs.h>
#include <linux/seq_file.h>

#define PROC_FILENAME "cpu_so1_1s2024"

static int cpu_proc_show(struct seq_file *m, void *v) {
    char buffer[256];
    struct file *fp;
    ssize_t n;

    // Abrir el archivo de /proc/stat
    fp = filp_open("/proc/stat", O_RDONLY, 0);

    // Leer la salida de /proc/stat
    n = kernel_read(fp, buffer, sizeof(buffer) - 1, &fp->f_pos);

    buffer[n] = '\0';

    // Escribir la salida en el archivo /proc
    seq_printf(m, "%s", buffer);
    filp_close(fp, NULL);
    
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
