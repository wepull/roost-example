all: pip_install jaeger_client opentracing_instrumentation booking
	echo "Check http://roost-utility:16686/"

pip_install:
	sudo apt-get -y install python-pip

jaeger_client:
	/usr/bin/pip install jaeger_client
	docker ps | grep 'jaegertracing/all-in-one:latest' | cut -f1 -d' ' | xargs -r docker stop -t0 2>/dev/null
	docker run --restart=always -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest

opentracing_instrumentation:
	/usr/bin/pip install opentracing_instrumentation

booking:
	/usr/bin/python booking-mgr.py 'A Beautiful Day in the Neighborhood'
	/usr/bin/python booking-mgr.py 'Dolittle'
	/usr/bin/python booking-mgr.py '1917'
	/usr/bin/python booking-mgr.py 'As Good As it Gets'
	/usr/bin/python booking-mgr.py 'Life is Beautiful'
	/usr/bin/python booking-mgr.py 'The Bridge on the River Kwai'