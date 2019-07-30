import java.io.IOException;
import java.util.Scanner;
public class Main {
 
    public static void main(String[] args) throws IOException {
 
		Scanner s= new Scanner(System.in);
		double moneda=0;
		int u100=0, u50=0, u20=0, u10=0, u5=0, u2=0, u1=0;
		int d50=0, d25=0, d10=0, d05=0, d01=0; 
		
		moneda=s.nextDouble();

		if (moneda>=100){
			u100=(int) (moneda/100);
			moneda=moneda-(u100*100);
		}
	
		if (moneda>=50){
			u50=(int) (moneda/50);
			moneda=moneda-(u50*50);
		}
	
		if (moneda>=20){
			u20=(int) (moneda/20);
			moneda=moneda-(u20*20);
		}
		
		if (moneda>=10){
			u10=(int) (moneda/10);
			moneda=moneda-(u10*10);
		}
		
		if (moneda>=5){
			u5=(int) (moneda/5);
			moneda=moneda-(u5*5);
		}
	
		if (moneda>=2){
			u2=(int) (moneda/2);
			moneda=moneda-(u2*2);
		}
		
		if (moneda>=1){
			u1=(int) (moneda/1);
			moneda=moneda-(u1*1);
		}
		
		if (moneda>=0.50){
			d50=(int) (moneda/0.50);
			moneda=(double) (moneda-(d50*0.50));
		}
	
		if (moneda>=0.25){
			d25=(int) (moneda/.25);
			moneda=(double) (moneda-(d25*.25));
		}
	
		if (moneda>=0.10){
			d10=(int) (moneda/.10);
			moneda=(double) (moneda-(d10*.10));
		}
	
		if (moneda>=0.05){
			d05=(int) (moneda/.05);
			moneda=(double) (moneda-(d05*.05));
		}
	
		if (moneda>=0.01){
			d01=(int) (moneda/0.01);
			moneda=(double) (moneda-(d01*0.01));
		}
	
		System.out.println("NOTAS:");
		System.out.println(u100 + " nota(s) de R$ 100.00");
		System.out.println(u50 + " nota(s) de R$ 50.00");	
		System.out.println(u20 + " nota(s) de R$ 20.00");
		System.out.println(u10 + " nota(s) de R$ 10.00");
		System.out.println(u5 + " nota(s) de R$ 5.00");
		System.out.println(u2 + " nota(s) de R$ 2.00");
		System.out.println("MOEDAS:");
		System.out.println(u1 + " moeda(s) de R$ 1.00");
		System.out.println(d50 + " moeda(s) de R$ 0.50");
		System.out.println(d25 + " moeda(s) de R$ 0.25");
		System.out.println(d10 + " moeda(s) de R$ 0.10");
		System.out.println(d05 + " moeda(s) de R$ 0.05");
		System.out.println(d01 + " moeda(s) de R$ 0.01");
		s.close();
    }
 
}
