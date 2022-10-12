class BirdWatcher {

    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return new int[] {0, 2, 5, 3, 7, 8, 4};
    }

    public int getToday() {
        if (this.birdsPerDay.length == 0) {
            return 0;
        }

        return this.birdsPerDay[this.birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        int newCount = this.getToday() + 1;

        this.birdsPerDay[this.birdsPerDay.length - 1] = newCount;
    }

    public boolean hasDayWithoutBirds() {
        for (int birdsCount : this.birdsPerDay) {
            if (birdsCount == 0) {
                return true;
            }
        }

        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        if (this.birdsPerDay.length == 0) {
            return 0;
        }

        int count = 0;

        if (numberOfDays > this.birdsPerDay.length) {
            numberOfDays = this.birdsPerDay.length;
        }

        for (int i = 0; i < numberOfDays; i++) {
            count += this.birdsPerDay[i];
        }

        return count;
    }

    public int getBusyDays() {
        int count = 0;

        for (int birdsCount : this.birdsPerDay) {
            if (birdsCount >= 5) {
                count++;
            }
        }

        return count;
    }

}
