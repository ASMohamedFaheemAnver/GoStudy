public class fibonacci {
    
    static int fibonacci(int n) {
        if (n <= 1) {
            return n;
        }
        int a = 0, b = 1;
        for (int i = 2; i <= n; i++) {
            int temp = b;
            b = a + b;
            a = temp;
        }
        return b;
    }
    
    public static void main(String[] args) {
        int n = 4096;
        long start = System.nanoTime();
        int result = fibonacci(n);
        long elapsed = System.nanoTime() - start;
        
        System.out.printf("Fibonacci(%d) = %d\n", n, result);
        System.out.printf("Time taken: %.6f ms\n", elapsed / 1_000_000.0);
    }
}
