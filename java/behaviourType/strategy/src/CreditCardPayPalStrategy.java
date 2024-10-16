class CreditCardPayPalStrategy implements Strategy {
    private String cardNumber;

    public CreditCardPayPalStrategy(String cardNumber) {
        this.cardNumber = cardNumber;
    }

    @Override
    public void pay(int amount) {
        System.out.println("product " + amount + " using Credit Card: " + cardNumber);
    }
}