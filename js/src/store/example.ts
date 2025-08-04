export function Store() {
  window.Alpine.data("counter", () => ({
    count: 0,
    get text() {
      return `TIMES CLICKED: ${this.count}`;
    },
    increment() {
      this.count++;
    },
    decrement() {
      this.count--;
    },
  }));
}
