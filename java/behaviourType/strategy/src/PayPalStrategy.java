public class PayPalStrategy implements Strategy {
    private String email;

    public PayPalStrategy(String email) {
        this.email = email;
    }

    @Override
    public void pay(int amount) {
        System.out.println("product " + amount + " using PayPal: " + email);
    }
}

