
# ==================================================================================== # # PRODUCTION
# ==================================================================================== #
production_host_ip = '209.38.16.135'
## production/connect: connect to the production server
.PHONY: production/connect 
production/connect:
	ssh ljpurcell@${production_host_ip}
