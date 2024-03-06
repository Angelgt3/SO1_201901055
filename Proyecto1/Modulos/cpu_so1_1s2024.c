#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>

static int __init cpu_module_init(void) {
    struct task_struct *task;
    int count = 0;
    // Iterar sobre todos los procesos para contarlos
    for_each_process(task) {
        count++;
    }
    // Imprimir el total de procesos en el kernel log
    printk(KERN_INFO "Total number of processes: %d\n", count);
    return 0; // Retornar 0 indica que la inicialización fue exitosa
}

static void __exit cpu_exit(void) {
    printk(KERN_INFO "Exiting cpu_so1_1s2024 module\n");
}

module_init(cpu_module_init);
module_exit(cpu_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel");
MODULE_DESCRIPTION("Módulo del kernel para obtener información de la CPU y los procesos");
