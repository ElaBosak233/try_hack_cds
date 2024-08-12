package model

type Submission struct {
	ID              uint           `json:"id"`                                                                                   // The submission's id. As primary key.
	Flag            string         `gorm:"type:varchar(128);not null" json:"flag,omitempty"`                                     // The flag which was submitted for judgement.
	Status          int            `gorm:"not null;default:0" json:"status"`                                                     // The status of the submission. (0-meaningless, 1-accepted, 2-incorrect, 3-cheat, 4-invalid(duplicate, etc.))
	UserID          uint           `gorm:"not null" json:"user_id"`                                                              // The user who submitted the flag.
	User            *User          `json:"user"`                                                                                 // The user who submitted the flag.
	ChallengeID     uint           `gorm:"not null;" json:"challenge_id"`                                                        // The challenge which is related to this submission.
	Challenge       *Challenge     `json:"challenge"`                                                                            // The challenge which is related to this submission.
	GameChallengeID *uint          `gorm:"index;null;default:null" json:"game_challenge_id,omitempty"`                           // The game_challenge which is related to this submission.
	GameChallenge   *GameChallenge `gorm:"foreignkey:GameChallengeID;association_foreignkey:ID" json:"game_challenge,omitempty"` // The game_challenge which is related to this submission.
	TeamID          *uint          `gorm:"index;null;default:null" json:"team_id,omitempty"`                                     // The team which submitted the flag. (Must be set when GameID is set)
	Team            *Team          `gorm:"foreignkey:TeamID;association_foreignkey:ID" json:"team,omitempty"`                    // The team which submitted the flag.
	GameID          *uint          `gorm:"index;null;default:null" json:"game_id,omitempty"`                                     // The game which is related to this submission. (Must be set when TeamID is set)
	Game            *Game          `gorm:"foreignkey:GameID;association_foreignkey:ID" json:"game,omitempty"`                    // The game which is related to this submission.
	Rank            int64          `json:"rank"`                                                                                 // The rank of the submission.
	Pts             int64          `gorm:"-" json:"pts"`                                                                         // The points of the submission.
	CreatedAt       int64          `gorm:"autoUpdateTime:milli" json:"created_at,omitempty"`                                     // The submission's creation time.
	UpdatedAt       int64          `gorm:"autoUpdateTime:milli" json:"updated_at,omitempty"`                                     // The submission's last update time.
}

func (s *Submission) Simplify() {
	s.Flag = ""
}
