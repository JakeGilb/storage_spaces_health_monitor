# Storage Spaces Health Monitor

This project provides a monitoring solution for Windows Storage Spaces running in a Kubernetes environment. It continuously checks the health of your storage pools and virtual disks, reporting any detected errors or issues to help ensure data integrity and system reliability.

## Features
- Monitors Windows Storage Spaces health status
- Integrates with Kubernetes for automated deployment and management
- Reports errors and warnings for pools and virtual disks
- Can be extended to send alerts via email, Slack, or other notification systems

## How It Works
The monitor runs as a containerized application within your Kubernetes cluster. It queries Windows Storage Spaces health using PowerShell commands and exposes the results for aggregation or alerting.

## Getting Started
1. **Clone the repository:**
   ```sh
   git clone https://github.com/JakeGilb/storage_spaces_health_monitor.git
   ```
2. **Build and deploy the container:**
   - Create a Docker image for the monitor
   - Deploy it as a Kubernetes Pod or DaemonSet
3. **Configure monitoring and alerting:**
   - Set up Prometheus scraping, or configure your preferred alerting system

## Example PowerShell Health Check
The monitor uses PowerShell commands like:
```powershell
Get-StoragePool | Get-StorageHealthReport
Get-VirtualDisk | Get-StorageHealthReport
```

## Customization
- Extend the monitor to support custom alerting endpoints
- Adjust the polling interval and error thresholds as needed

## Contributing
Pull requests and suggestions are welcome! Please open an issue to discuss your ideas or report bugs.

## License
MIT
