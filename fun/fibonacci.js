function fibonacci(n) {
  if (n <= 1) return n;

  let a = 0,
    b = 1;
  for (let i = 2; i <= n; i++) {
    [a, b] = [b, a + b];
  }
  return b;
}

const n = 4096;

const start = process.hrtime.bigint();
const result = fibonacci(n);
const end = process.hrtime.bigint();

const timeTakenMs = Number(end - start) / 1_000_000;

console.log(`Fibonacci(${n}) = ${result}`);
console.log(`Time taken: ${timeTakenMs.toFixed(3)} ms`);
