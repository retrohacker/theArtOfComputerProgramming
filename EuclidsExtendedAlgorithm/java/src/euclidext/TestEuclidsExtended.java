package euclidext;

import org.junit.*;
import junit.framework.TestCase;

/**
 * Unit tests for the EuclidExtended class.
 */
public class TestEuclidsExtended extends TestCase {

	//Table based unit test
	Result[] testCases = {
		new Result(1769,551,5,-16,29,false,null),
		new Result(551,1769,-16,5,29,false,null),
		new Result(-1769,551,0,0,0,true,EuclidsExtended.NonPos),
		new Result(120,23,-9,47,1,false,null),
		new Result(-120,23,0,0,0,true,EuclidsExtended.NonPos),
		new Result(23,120,47,-9,1,false,null),
		new Result(46,57,-26,21,1,false,null),
		new Result(46,-57,0,0,0,true,EuclidsExtended.NonPos)
	};

	/**
	 * Table based batch test of the function.
	 */
	@Test
	public void testFunction() {
		for(int i=0; i<testCases.length;i++) {
			EuclidsExtended testClass = new EuclidsExtended(testCases[i].input1,testCases[i].input2);
			testClass.execute();
			assertTrue("Expected: "+testCases[i]+" Result:"+testClass.getResult(),testCases[i].equals(testClass.getResult()));
		}
	}
}
