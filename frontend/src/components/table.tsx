import { Matches } from "../models/matches.interface";
import '../css/table.css';

interface TableProps {
    matchData: Matches;
}

export default function CustomizedTables({ matchData }: TableProps): JSX.Element {
    const shortenMatchList = matchData.match_list.slice(0, 5); 

    return (
        <table>
            <thead>
                <tr>
                    <th key={matchData.tournament_id} colSpan={2}>{matchData.game_name}</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th key={"match"}>Match</th>
                    <th key={"round"}>Round</th>
                </tr>
                
            </thead>
            <tbody>
                {shortenMatchList.map(match => (
                    <tr>
                        <td key={Math.random()}>{`${match.player1_name} vs ${match.player2_name}`}</td>
                        <td key={Math.random()}>{match.round <= -1 ? `losers ${Math.abs(match.round)}`: `winners ${match.round}`}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}
