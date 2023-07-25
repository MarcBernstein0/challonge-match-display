import * as React from 'react';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { Match, Matches } from '../models/matches.interface';

interface Column {
  id: 'match' | 'round' | 'underway';
  label: string;
  minWidth?: number;
  align?: 'right' | 'center';
  format?: (value: boolean) => string;
}

const columns: readonly Column[] = [
  { id: 'match', label: 'Match', align: 'center', minWidth: 170 },
  { id: 'round', label: 'Round', align: 'center', minWidth: 100 },
  {
    id: 'underway',
    label: 'Underway',
    minWidth: 170,
    align: 'center',
  },
];

interface TableProps {
  matchData: Matches;
}

interface Data {
  match: string;
  round: string;
  underway: boolean;
}

function createData(
  matchData: Match
): Data {
  const match = `${matchData.player1_name} vs ${matchData.player2_name}`;
  const round = matchData.round <= -1 ? `winners ${Math.abs(matchData.round)}` : `losers ${Math.abs(matchData.round)}`;
  const underway = matchData.underway;
  return { match, round, underway };
}

function createRows(matchList: Match[]): Data[] {
  const res: Data[] = [];
  for (const match of matchList) {
    res.push(createData(match));
  }
  return res;
}

export default function StickyHeadTable({ matchData }: TableProps) {
  const matches: Match[] = matchData.match_list;
  const rows: Data[] = createRows(matches);

  return (
    <Paper sx={{ 
      width: '100%', 
      overflow: 'hidden',
      borderStyle: 'solid',
      borderColor: 'black' }}>
      <TableContainer sx={{ 
        maxHeight: 440, 
        borderStyle: 'solid',
        borderColor: 'black'}}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              <TableCell align="center" colSpan={3}>
                {matchData.game_name}
              </TableCell>
            </TableRow>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ top: 57, minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows
              .map((match) => {
                return (
                  <TableRow hover key={match.match}>
                    <TableCell align='center'>
                      {match.match}
                    </TableCell>
                    <TableCell align='center'>
                      {match.round}
                    </TableCell>
                    <TableCell align='center'
                      sx={{
                        color: match.underway ? 'green' : 'red',
                        fontSize: '1.2vw'
                      }}>
                      {match.underway ? `Yes` : `No`}
                    </TableCell>
                  </TableRow>
                );
              })}
          </TableBody>
        </Table>
      </TableContainer>
    </Paper>
  );
}