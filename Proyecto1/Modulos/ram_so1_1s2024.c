#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/mm.h>
#include <linux/seq_file.h>
#include <linux/fs.h>

#define PROC_FILENAME "ram_so1_1s2024"

static int ram_proc_show(struct seq_file *m, void *v) {
    struct sysinfo info;
    si_meminfo(&info);
    seq_printf(m, "Total RAM: %lu\n", info.totalram);
    seq_printf(m, "Free RAM: %lu\n", info.freeram);
    return 0;
}

static int ram_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, ram_proc_show, NULL);
}

static const struct proc_ops ram_proc_fops = {
    .proc_open   = ram_proc_open,
    .proc_read   = seq_read,
    .proc_lseek  = seq_lseek,
    .proc_release= single_release,
};

static int __init ram_init(void) {
    proc_create(PROC_FILENAME, 0, NULL, &ram_proc_fops);
    printk(KERN_INFO "ram_so1_1s2024 module initialized\n");
    return 0;
}

static void __exit ram_exit(void) {
    remove_proc_entry(PROC_FILENAME, NULL);
    printk(KERN_INFO "Exiting ram_so1_1s2024 module\n");
}

module_init(ram_init);
module_exit(ram_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel");
MODULE_DESCRIPTION("A simple RAM info kernel module");
