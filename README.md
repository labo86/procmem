# procmem
Programa para medir el uso de memoria

## Instalar
Bájelo con el siguiente comando
```wget https://github.com/imo-chile/procmem/releases/download/v0.1/procmem_linux -O procmem```
o usando el siguiente [link](https://github.com/imo-chile/procmem/releases/download/v0.1/procmem_linux).

## Uso
`./procmem output interval comando`
 - __output__: el nombre del archivo de salida
 - __interval__:  el intervalo de muestreo de memoria en segundos
 - __comando__: el comando a medir
 
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

La salida tiene 3 secciones, __START__, __RUNNING__, __FINISHED__.
### START
En la sección __START__ se muestra la información de la ejecución.

### RUNNING
En la sección __RUNNING__ está registrado el muestro periódico del uso de memoria. La frecuencia esta dada por la duración del intervalo.
Los valores __VmSize__, __VmPeak__, __VmRSS__ y __VmHWN__ son extraídos de `/proc/$PID/stat` donde PID es el número de proceso. Para más información ver [man proc](https://man7.org/linux/man-pages/man5/proc.5.html).

### FINISHED
Los valores corresponden al tiempo de usuario y el tiempo de sistema. El __UTime__ es extraído con la función [ProcessState.UserTime()](https://golang.org/pkg/os/#ProcessState.UserTime) y el __STime__ con [ProcessState.SystemTime()](https://golang.org/pkg/os/#ProcessState.SystemTime).


