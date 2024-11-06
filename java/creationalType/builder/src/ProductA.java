public class ProductA extends Product{
    /**
     * 产品A
     * 产品A的构建逻辑和构建顺序交由builder进行控制
     */

    public String name;
    public String radio;
    public String speaker;

    ProductA() {
        this.name = "product A";
    }

    public void setRadio(String radio) {
        this.radio = radio;
    }

    public void setSpeaker(String speaker) {
        this.speaker = speaker;
    }

    @Override
    void describe() {
        System.out.printf("it is product A ,my radio is %s , my speaker is %s \n", radio,speaker);
    }
}
