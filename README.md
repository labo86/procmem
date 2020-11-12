# procmem
Programa para medir el uso de memoria

## Instalar
Bájelo con el siguiente comando
`wget https://github.com/imo-chile/procmem/releases/download/v0.1/procmem_linux -O procmem`

## Uso
`./procmem output interval comando`
 - output: el nombre del archivo de salida
 - interval:  el intervalo de muestreo de memoria en segundos
 - comando: el comando a medir
 
## Ejemplo de salida
```
##START
Command : diamond
Args: [blastx -d ../../jcarcamo/mg/sprot -q ../../jcarcamo/mg/AA_man_anterior_nare1.fna -o result_0.txt --max-target-seqs 1 -p 40 --sensitive]
Interval: 15
##RUNNING
secs    VmSize  VmPeak  VmRSS   VmHWN
15      3881596 kB      3881596 kB      862316 kB       863984 kB
30      3611128 kB      3881596 kB      862216 kB       863984 kB
45      3881596 kB      3881596 kB      862684 kB       863984 kB
##FINISHED
UTime : 1639.412000
STime : 7.376000
```

La salida tiene 3 secciones, START, RUNNING, FINISHED.
### START
En la sección START se muestra la información de la ejecución.

### RUNNING
En la sección RUNNING se hace un muestreo periódico del uso de memoria. La frecuencia esta dada por la duración del intervalo.
Los valores VmSize VmPeak VmRSS y VmHWN son extraidos de /proc/$PID/stat donde PID es el número de proceso. Para más información ver [man proc](https://man7.org/linux/man-pages/man5/proc.5.html)


