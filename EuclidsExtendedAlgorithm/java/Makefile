JUNIT_HOME = /usr/share/java/junit4.jar
test:
	@rm -f bin/main/*.class
	@echo Compiling...
	@javac -cp "$(JUNIT_HOME)" -d bin src/euclidext/*
	@java -cp "bin:$(JUNIT_HOME)" org.junit.runner.JUnitCore euclidext.TestEuclidsExtended
	@echo Generating Javadoc
	@javadoc -private -sourcepath src/ -verbose -d doc/ euclidext
