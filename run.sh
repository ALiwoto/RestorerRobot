
buildApp() 
{
	# clear the screen (the terminal)
	clear

	echo -e "building wotoubBinary, please wait a bit..."

	go build -o wotoubBinary
}

runApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are done building it,\n->now running the wotoubBinary...\n-------------------"

	./wotoubBinary
}

testApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are running all test files (*_test.go);\nplease wait a bit"

	go test -v ./...
}

if [ "$1" == "test" ];
then
	testApp;
	exit 0
fi;

operations=0

if [ -z "$1" ] || [ "$1" == "true" ] || [ "$1" == "1" ];
then
	buildApp;
	operations=$((i+1))
fi;

if [ -z "$2" ] || [ "$2" == "true" ] || [ "$2" == "1" ];
then
	runApp;
	operations=$((i+1))
fi;

if [ $operations == 0 ]
then
	echo "You have done nothing!"
fi;