// START1 OMIT
class VisibilityTest {
	public int x;
	private int y;
	// END1 OMIT

	public static VisibilityTest factory() {
		// START2 OMIT
		VisibilityTest vt = new VisibilityTest();
		vt.x = 5;
		vt.y = 2;	
		// END2 OMIT
	}
}
