public class Main {
    public static void main(String[] args) {
        // context 上下文 可以理解为执行器
        Context context = new Context();

        // 具体的执行策略
        Strategy payPalPayment = new PayPalStrategy("");
        Strategy creditCardPayPalStrategy = new CreditCardPayPalStrategy("");

        // 策略执行 然后切换上下文
        context.setPaymentStrategy(payPalPayment);
        context.pay(100);

        context.setPaymentStrategy(creditCardPayPalStrategy);
        context.pay(100);

    }
}