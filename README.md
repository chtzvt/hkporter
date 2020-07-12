# Porter HomeKit Integration`

This repository contains a HomeKit integration proxy for [porter](https://github.com/ctrezevant/porter), a smart garage door controller.

Configuration options (passed as command line flags):
```
-pin     HomeKit pairing PIN"
-dbpath  State database path (optional, default is ./db)
-v       Enable HomeKit server debug output")
-api     Porter API server URI (default: http://localhost:80)
-key     Porter API key
```

This repository contains a Systemd unit file (hkporter.service) that can be used to run and manage this service.
