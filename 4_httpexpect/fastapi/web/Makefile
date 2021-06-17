# For development
start:
	uvicorn main:app --reload

# Take advantage of multicore
workers:
	uvicorn --workers 8 main:app

# All for more config (port for example)
py:
	python main.py