abstract class Fighter {

    boolean isVulnerable() {
        return false;
    }

    abstract int damagePoints(Fighter fighter);

}

class Warrior extends Fighter {

    @Override
    public String toString() {
        return "Fighter is a Warrior";
    }

    @Override
    int damagePoints(Fighter fighter) {
        return fighter.isVulnerable() ? 10 : 6;
    }
}

class Wizard extends Fighter {
    private boolean isSpellPrepared = false;

    @Override
    public String toString() {
        return "Fighter is a Wizard";
    }

    @Override
    boolean isVulnerable() {
        return !this.isSpellPrepared;
    }

    @Override
    int damagePoints(Fighter fighter) {
        return this.isSpellPrepared ? 12 : 3;
    }

    void prepareSpell() {
        this.isSpellPrepared = true;
    }

}
