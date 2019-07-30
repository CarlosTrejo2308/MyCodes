import java.io.IOException;
import java.text.DecimalFormat;
import java.util.Scanner;
public class Main {
 
    public static void main(String[] args) throws IOException {

	    Scanner s= new Scanner(System.in);
	    
	    double A,R;
		final double pi=3.14159;
        R=s.nextDouble();
        A=R*R;
        A=A*pi;
        
        DecimalFormat numberFormat = new DecimalFormat("0.0000");
        System.out.println("A=" + numberFormat.format(A));
s.close();

    }
 
}
