import { Match, Matches } from "../models/matches.interface";
import '../css/table.css';

interface TableProps {
  matchData: Matches;
}

export default function ArrayOfCustomizable( matchData: Matches): JSX.Element[] {
  const chunk = 5;
  const result: JSX.Element[] = [];
  const sliceOfMatchData = [];
  
  let currentVal = 0;
  while (currentVal < matchData.match_list.length) {
    sliceOfMatchData.push(matchData.match_list.slice(currentVal, currentVal + chunk));
    currentVal += chunk;
  }
  console.log(sliceOfMatchData);
  for(const sliceData of sliceOfMatchData) {
    result.push(CustomizedTables(matchData.tournament_id, matchData.game_name, sliceData));
  }

  return result;
}

function CustomizedTables(tournamentId: number, gameName: string, matchList: Match[]): JSX.Element {
  return (
    <table>
      <thead>
        <tr>
          <th key={tournamentId} colSpan={3}>{gameName}</th>
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
        {matchList.map(match => (
          <tr style={{
            fontWeight: 'bold',
          }}>
            <td key={Math.random()}>{`${match.player1_name} vs ${match.player2_name}`}</td>
            <td key={Math.random()}>{match.round <= -1 ? `losers ${Math.abs(match.round)}` : `winners ${match.round}`}</td>
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
