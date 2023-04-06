Benchmark in Go for binary search, bubble sort and reading file contents using Intel Power Gadget

I tested this on my Windows Intel(x64) laptop so you might need to investigate tools for other computer hardware architectures and Operating Systems:
Intel Power Gadget -> https://www.intel.com/content/www/us/en/developer/articles/tool/power-gadget.html

A slight caveat from my exeperience with the benchmarking tool (I was running in powershell), you need to pass '&' before execution of command (this is just a powershell that I didn't know that might catch you out)

sample command -> '& "C:\Program Files\Intel\Power Gadget 3.6\PowerLog3.0.exe" -file {file_name_to_save_to}.csv -verbose -cmd {command_to_run}'

With this command energy consumption report files are stored in directory of execution.

Go execution
& "C:\Program Files\Intel\Power Gadget 3.6\PowerLog3.0.exe" -file {file_name_to_save_to}.csv -verbose -cmd go test -benchmem -run=^$ -bench ^{benchmark_test_func_name}$ github/sekarasiewicz/go-python-benchmark/{directory_name}