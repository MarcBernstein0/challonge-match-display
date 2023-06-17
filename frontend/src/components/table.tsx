import { Match, Matches } from "../models/matches.interface";
import '../css/table.css';

interface TableProps {
    matchData: Matches;
}

export default function CustomizedTables({ matchData }: TableProps): JSX.Element {
    const matchDataList: Match[] = matchData.match_list;
    return (
        <table>
            <thead>
                <tr>
                    <th key={matchData.tournament_id} colSpan={3}>{matchData.game_name}</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th key={"match"}>Match</th>
                    <th key={"round"}>Round</th>
                    <th key={"underway"}>Underway</th>
                </tr>
                
            </thead>
            <tbody>
                {matchDataList.slice(0, 16).map(match => (
                    <tr style={{
                        fontWeight: 'bold',
                    }}>
                        <td key={Math.random()}>{`${match.player1_name} vs ${match.player2_name}`}</td>
                        <td key={Math.random()}>{match.round <= -1 ? `losers ${Math.abs(match.round)}`: `winners ${match.round}`}</td>
                        <td key={Math.random()}
                            style={{
                                color: match.underway ? 'green' : 'red',
                                fontSize: '1.2vw'
                            }}
                            >{match.underway ? `Yes` : `No`}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}
