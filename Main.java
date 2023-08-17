package edualb.aubay_interview;

public class Main {
    /*
    * War of Numbers
    * 
    * There's a great war between the even and odd numbers. 
    * Many numbers already lost their lives in this war and it's your task to end this. 
    * You have to determine which group sums larger: the evens, or the odds. 
    * The larger group wins.
    * 
    * Create a function that takes an array of integers, sums the even and odd numbers separately, 
    * then returns the difference between the sum of the even and odd numbers.
    * 
    * Examples
    * warOfNumbers([2, 8, 7, 5]) ➞ 2
    * warOfNumbers([12, 90, 75]) ➞ 27
    * warOfNumbers([5, 9, 45, 6, 2, 7, 34, 8, 6, 90, 5, 243]) ➞ 168
    * 
    * Source: https://edabit.com/challenge/7fHsizQrTLXsPWMyH
    */
    public static void main(String[] args) {
        int[] input = {2, 8, 7, 5};
        // int[] input = {12, 90, 75};
        // int[] input = {5, 9, 45, 6, 2, 7, 34, 8, 6, 90, 5, 243};

        int output = warOfNumbers(input);
        System.out.println("Output: " + output);
    }

    public static Integer warOfNumbers(int[] in) {
        /* your logic here */
        return 0;
    }
}