#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/sysinfo.h>
#include <linux/mm.h>

static int __init ram_init(void) {
    struct sysinfo info;
    si_meminfo(&info);
    printk(KERN_INFO "Total RAM: %lu\n", info.totalram);
    printk(KERN_INFO "Free RAM: %lu\n", info.freeram);
    return 0;
}

static void __exit ram_exit(void) {
    printk(KERN_INFO "Exiting ram_so1_1s2024 module\n");
}

module_init(ram_init);
module_exit(ram_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel");
MODULE_DESCRIPTION("A simple RAM info kernel module");
