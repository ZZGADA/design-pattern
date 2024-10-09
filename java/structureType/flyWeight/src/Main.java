public class Main {
    public static void main(String[] args) {
        FlyWeightFactory flyWeightFactory = FlyWeightFactory.GetFlyWeightFactory();

        FlyWeight firstFlyWeight = flyWeightFactory.get("ConcreteFlyWeightFirst");
        FlyWeight secondFlyWeight = flyWeightFactory.get("ConcreteFlyWeightSecond");

        FlyWeight[] flyWeights = {firstFlyWeight, secondFlyWeight};
        for(FlyWeight flyWeight : flyWeights) {
            flyWeight.say();
        }

    }
}