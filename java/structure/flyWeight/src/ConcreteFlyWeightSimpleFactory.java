public class ConcreteFlyWeightSimpleFactory {
    private static ConcreteFlyWeightSimpleFactory instance;

    public static  ConcreteFlyWeightSimpleFactory GetFactory() {
        if (instance == null) {
            synchronized (ConcreteFlyWeightSimpleFactory.class) {
                if (instance == null) {
                    instance = new ConcreteFlyWeightSimpleFactory();
                }
            }
        }
        return instance;
    }


    public FlyWeight GetConcreteFlyWeight(String key) {
        switch (key) {
            case "ConcreteFlyWeightFirst":
                return new ConcreteFlyWeightFirst();
            case "ConcreteFlyWeightSecond":
                return new ConcreteFlyWeightSecond();
            default:
                return null;
        }
    }

}
