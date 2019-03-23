import java.util.*; 

public class Scoping1 {


    public static void main(String [] args){
       int inner; // notice that when we uncomment this line and remove the int word on line 7 the program will compile.
        if (true) { // demonstration of block level scoping
            inner = 0;
            { // scoping braces
                int locally = 1;
                System.out.println(locally);
            }
            System.out.println(inner); // the program will fail to compile here if we changed inner to locally
        }
        System.out.println(inner); // the program will fail to compile here if we changed inner to locally

    }

}