This is a Python example for Jaeger Open Tracing

1) Install the jaeger-client
   pip install jaeger-client

2) # Check if docker is up and running
   docker ps 

3) Run the jaegertracing image
   docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest

4) Once the container starts, check the UI. The container runs the Jaeger backend with an in-memory store, which is initially empty.
   Open http://localhost:16686/ in a browser

5) Creating Traces on Jaeger UI

   a) pip install opentracing_instrumentation

   #Run the python program to look for a movie name and create a booking
   b) python booking-mgr.py <movie-name>
   

Sample output:
harishagrawal@Harishs-MacBook-Pro opentracing % python booking-mgr.py 'abc'            
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

