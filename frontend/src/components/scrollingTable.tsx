import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Match, Matches } from '../models/matches.interface';

interface TableProps {
  matchData: Matches;
}

export default function ScrollingTable({ matchData }: TableProps): JSX.Element {
  const matchDataList: Match[] = matchData.match_list;
  return (
    <TableContainer component={Paper}>
      <Table sx={{ 
        minWidth: 550,
       }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell colSpan={3} align="center" sx={{ fontWeight: 'bold' }}>{matchData.game_name}</TableCell>
          </TableRow>
        </TableHead>
        <TableHead>
          <TableRow>
            <TableCell align="center" sx={{ fontWeight: 'bold' }}>Match</TableCell>
            <TableCell align="center" sx={{ fontWeight: 'bold' }}>Round</TableCell>
            <TableCell align="center" sx={{ fontWeight: 'bold' }}>Underway</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {matchDataList.map((match) => (
            <TableRow
              key={match.player1_name}
              sx={{
                
              }}
            >
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {`${match.player1_name} vs ${match.player2_name}`}
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>{match.round <= -1 ? `losers ${Math.abs(match.round)}`: `winners ${match.round}`}</TableCell>
              <TableCell align="center" sx={{
                color: match.underway ? 'green' : 'red',
                fontSize: '1.2vw',
                fontWeight: 'bold'
              }}>{match.underway ? `Yes` : `No`}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}