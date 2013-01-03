JUNIT_HOME = "/usr/share/java/junit4.jar"
test :
	@rm -f main/*.class
	@echo Compiling...
	@javac -cp $(JUNIT_HOME) main/*
	@java -cp "./:$(JUNIT_HOME)" org.junit.runner.JUnitCore main.TestEuclid
