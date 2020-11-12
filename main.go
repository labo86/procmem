package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type MemStat struct {
	VmSize string
	VmPeak string
	VmHWM string
	VmRSS string
}

func ScanMemStat(r io.Reader) *MemStat {
	var mem MemStat

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Vm") {
			values := strings.Split(line, ":")
			name := strings.TrimSpace(values[0])
			value:= strings.TrimSpace(values[1])
			switch name {
			case "VmSize":
				mem.VmSize = value
			case "VmRSS":
				mem.VmRSS = value
			case "VmPeak":
				mem.VmPeak = value
			case "VmHWM":
				mem.VmHWM = value
			}
		}

	}
	return &mem

}

// https://man7.org/linux/man-pages/man5/proc.5.html
func getInfo(process *os.Process) *MemStat {

	filename := fmt.Sprintf("/proc/%d/status", process.Pid)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return ScanMemStat(file)
}

func Measure(command string, args []string, interval time.Duration, output io.Writer) error {
	fmt.Fprintln(output, "##START")
	fmt.Fprintf(output, "Command : %s\n", command)
	fmt.Fprintf(output, "Args: %+v\n", args)
	fmt.Fprintf(output, "Interval: %d\n", int(interval.Seconds()))


	startTime := time.Now()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	done := make(chan error)

	cmd := exec.Command(command, args...)
	if err := cmd.Start() ; err != nil {
		return fmt.Errorf("no se pudo correr el programa : %v", err)
	}

	go func() {
		done <- cmd.Wait()
	}()


	fmt.Fprintln(output, "##RUNNING")
	fmt.Fprintln(output, "secs\tVmSize\tVmPeak\tVmRSS\tVmHWN")
	for {
		select {
		case err := <-done:
			if err != nil {
				return fmt.Errorf("error en ejecucion de programa : %v", err)
			}

			fmt.Fprintln(output, "##FINISHED")
			fmt.Fprintf(output, "UTime : %f\nSTime : %f\n", cmd.ProcessState.UserTime().Seconds(), cmd.ProcessState.SystemTime().Seconds())
			return nil
		case t := <-ticker.C:
			value  := getInfo(cmd.Process)
			seconds := int(t.Sub(startTime).Seconds())

			fmt.Fprintf(output, "%d\t%s\t%s\t%s\t%s\n", seconds, value.VmSize, value.VmPeak, value.VmRSS, value.VmHWM)
		}
	}
}

func main() {
	argc := len(os.Args)

	if argc < 4 {
		fmt.Printf("Uso : procmem salida interval(secs) comando\n")
		return
	}
	outputFile := os.Args[1]
	command := os.Args[3]

	interval, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Uso : procmem salida interval(secs) comando\n")
		return
	}

	if interval <=0 {
		fmt.Printf("Uso : procmem salida interval(secs) comando\n")
		return
	}

	var args []string
	if argc > 4 {
		args = os.Args[4:]
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no se puede abrir archivo de salida %q : %v\n", outputFile, err)
		os.Exit(1)
	}

	if  err := Measure(command, args, time.Duration(interval) * time.Second, output); err != nil {
		fmt.Fprintf(os.Stderr,"error al hacer la medición : %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)





}
