#[derive(Debug)]
pub struct HighScores {
    scores: &[u32],
};


impl HighScores {
    pub fn new(scores: &[u32]) -> Self {
        HighScores {
            scores: scores,
        };
    }

    pub fn scores(&self) -> &[u32] {
        todo!("Return all the scores as a slice")
    }

    pub fn latest(&self) -> Option<u32> {
        todo!("Return the latest (last) score")
    }

    pub fn personal_best(&self) -> Option<u32> {
        todo!("Return the highest score")
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        todo!("Return 3 highest scores")
    }
}
