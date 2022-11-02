import { Matches } from "../models/matches.interface";
import '../css/table.css';

interface TableProps {
    matchData: Matches;
}

export default function CustomizedTables({ matchData }: TableProps): JSX.Element {
    return (
        <table>
            <thead>
                <tr>
                    <th colSpan={2}>{matchData.game_name}</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th>Match</th>
                    <th>Round</th>
                </tr>
                
            </thead>
            <tbody>
                {matchData.match_list.map(match => (
                    <tr>
                        <td>{`${match.player1_name} vs ${match.player2_name}`}</td>
                        <td>{match.round}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}
