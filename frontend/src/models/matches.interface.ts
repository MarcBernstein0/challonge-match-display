export interface Match {
    id: number;
    player1_id: number;
    player1_name: string;
    player2_id: number;
    player2_name: string;
    round: number;
}

export interface Matches {
    game_name: string;
    tournament_id: number;
    match_list: Match[];
}
