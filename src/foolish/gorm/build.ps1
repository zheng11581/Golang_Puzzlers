Set-Culture zh-CN
$OutputEncoding = [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$InputEncoding = [Console]::InputEncoding = [System.Text.Encoding]::UTF8
# 指定 Go 项目目录
$projectPath = Get-Location
# 进入项目目录
Set-Location $projectPath
# 交叉编译为 Linux 平台的可执行文件
Write-Host "Start building Linux platform binary..."
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o $projectPath/toy-gorm

# 如果需要其他平台的交叉编译，可以按照上述格式进行添加
Write-Host "Build complete!"