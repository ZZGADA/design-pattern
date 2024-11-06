public class ProductB extends Product{
    /**
     * 产品B
     * 产品B的构建逻辑和构建顺序交由builder进行控制
     * 与产品A不同 二者可能产品类型一致 但是构建顺序 和具体的构建组成部件不通过
     */

    public String name;
    public String radio;
    public String speaker;

    ProductB() {
        this.name = "product B";
    }

    public void setRadio(String radio) {
        this.radio = radio;
    }

    public void setSpeaker(String speaker) {
        this.speaker = speaker;
    }



    @Override
    void describe() {
        System.out.printf("it is product b ,my radio is %s , my speaker is %s \n", radio,speaker);
    }

}
