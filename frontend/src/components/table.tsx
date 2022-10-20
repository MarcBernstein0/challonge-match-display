import { Matches } from "../models/matches.interface";
import '../css/table.css';

interface TableProps {
    matchData: Matches;
}

export default function CustomizedTables({ matchData }: TableProps): JSX.Element {
    // console.log("call in CustomizedTables component", matchData);
    // const game1 = matchData[0];
    // console.log(matchData);
    return (
        <table>
            <thead>
                <th colSpan={2}>{matchData.game_name}</th>
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
        // <table style={{"borderWidth":"1px", 'borderColor':"#aaaaaa", 'borderStyle':'solid'}}>
        //     <tr>
        //         <th>{matchData.game_name}</th>
        //     </tr>
        //     {matchData.match_list.map(match => (
        //         <tr>
        //             <td>{`${match.player1_name} vs ${match.player2_name}`}</td>
        //             <td>{match.round}</td>
        //         </tr>
        //     ))}

        // </table>

    );
}
