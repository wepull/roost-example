set -x
echo "Calling Sample producer"

echo $PWD
ls -l /zbsample/zbsample_producer

export SERVER_CRT=/zbsample/server.crt 
touch i_am_here.txt
PROG=$(ls -lart . | grep zbsample | awk '{print $NF}')
ldd $PROG
./$PROG

echo "Done"
