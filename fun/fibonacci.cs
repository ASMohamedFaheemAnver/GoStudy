using System;
using System.Diagnostics;

class fibonacci {
    
    static int Fibonacci(int n) {
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
    
    static void Main() {
        int n = 4096;
        Stopwatch sw = Stopwatch.StartNew();
        int result = Fibonacci(n);
        sw.Stop();
        
        Console.WriteLine($"Fibonacci({n}) = {result}");
        Console.WriteLine($"Time taken: {sw.Elapsed.TotalMilliseconds:F6} ms");
    }
}
