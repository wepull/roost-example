## Python program using Jaeger Open Tracing

## How to run this application

###### Using Roost Desktop Engine (RDE)

> Right-click on `Makefile` and click `Run` for hassle-free deployment to ZKE

 >What all is done by `make`?
  * Installs python-pip 
  * Installs jaeger_client
  * Installs opentracing_instrumentation
  * Executes python program booking-mgr.py
  

###### _Using RKT Konsole_
```bash
1) Install the jaeger-client
   sudo apt-get install python-pip
   pip install jaeger-client

2) Confirm that docker is up and running
   docker ps 

3) Run the jaegertracing image
   docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest

4) To access the trace/spans at http://roost-utility:16686/ in any browser.

5) Creating Traces on Jaeger UI

   a) pip install opentracing_instrumentation

   #Run the python program to look for a movie name and create a booking
   b) python booking-mgr.py <movie-name>
```   

###### Sample output:

```bash
$ python booking-mgr.py 'abc'      
      
Initializing Jaeger Tracer with UDP reporter
Using selector: KqueueSelector
Using sampler ConstSampler(True)
opentracing.tracer initialized to <jaeger_client.tracer.Tracer object at 0x10b01a050>[app_name=booking]
Reporting span b728a942aaf2ae76:c5125d1ec0e1b86f:21bf675e0d8009a2:1 booking.CheckCinema
Reporting span b728a942aaf2ae76:7a5a6a75baa1b3f6:21bf675e0d8009a2:1 booking.CheckShowtime
Ticket Details
Reporting span b728a942aaf2ae76:45bf0d82d30f1073:21bf675e0d8009a2:1 booking.BookShow
Reporting span b728a942aaf2ae76:21bf675e0d8009a2:0:1 booking.booking
Using selector: KqueueSelector
```

``` 
Raise any issue or feature request using RDE Help
Join the Awesome Roost Community https://join.slack.com/t/roostai/shared_invite/zt-ea5mo10y-jDJgXiHn0RihSmucz0UZpw
```
