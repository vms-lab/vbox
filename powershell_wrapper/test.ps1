# $vbox = New-Object -ComObject VirtualBox.VirtualBox
# $vbox | ConvertTo-Json

$MediumVariant = New-Object -TypeName VirtualBox.MediumVariant.MediumVariant_Standard
Write-Host $MediumVariant