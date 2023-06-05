package constant

const HOSTNAME = "hostname"

const CPU = `(Get-WmiObject -Query "select Name, PercentProcessorTime , PercentUserTime , PercentIdleTime from Win32_PerfFormattedData_PerfOS_Processor") | foreach-object { write-host "$($_.Name);$($_.PercentProcessorTime);$($_.PercentUserTime);$($_.PercentIdleTime)" };`

const PROCESS = `(Get-Process -IncludeUsername) | Select-Object | foreach-object { write-host "$($_.Id);$($_.WS);$($_.CPU);$($_.path);$($_.ProcessName)"};`

const DISK = `(Get-WmiObject -Query "select Name, Size , freespace , VolumeName from Win32_LogicalDisk") | foreach-object { write-host "$($_.Name);$($_.Size+0);$($_.freespace+0);$($_.Size - $_.freespace);$([Math]::Round(($_.freespace / ($_.size+1)) * 100, 2));$([Math]::Round(((($($_.Size - $_.freespace) / ($_.size+1)) * 100)), 2));$($_.VolumeName)" };`

const SYSTEMINFO = `$data0= hostname
$data1= (Get-WmiObject -Query "select  Name , Version from Win32_OperatingSystem");
$data2 = (Get-WmiObject -Query "select  Processes , Threads , ContextSwitchesPersec , SystemUpTime from Win32_PerfFormattedData_PerfOS_System");
echo "$data0;$($data1.Name);$($data1.Version);$($data2.Processes);$($data2.Threads);$($data2.ContextSwitchesPersec);$($data2.SystemUpTime)"`

const MEMORY = `$data2 = (Get-WmiObject -Query "select  TotalVisibleMemorySize ,  FreePhysicalMemory , TotalSwapSpaceSize , TotalVirtualMemorySize  from CIM_OperatingSystem");
$totalMem = $data2.TotalVisibleMemorySize
$freeMem = $($data2.FreePhysicalMemory);
$usedMem = $($($data2.TotalVisibleMemorySize) - $($data2.FreePhysicalMemory))
$verMem = $data2.TotalVirtualMemorySize
$swapMem =  ($verMem - $totalMem)
$freePerc = $([Math]::Round((($freeMem/$totalMem)*100),2))
$usedPerc = $([Math]::Round((($usedMem/$totalMem)*100),2))
echo "$totalMem;$freeMem;$($usedMem);$freePerc;$usedPerc;$swapMem"`
