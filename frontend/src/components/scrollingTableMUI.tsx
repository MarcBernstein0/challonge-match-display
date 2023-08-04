import * as React from 'react';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import { Matches } from '../models/matches.interface';

interface Column {
  id: 'match' | 'round' | 'underway';
  label: string;
  minWidth?: number;
  align?: 'left' | 'center' | 'right';
  format?: (value: number) => string;
}

const columns: readonly Column[] = [
  { id: 'match', label: 'Match', minWidth: 50, align: 'center' },
  { id: 'round', label: 'Round', minWidth: 50, align: 'center' },
  {
    id: 'underway',
    label: 'Underway',
    minWidth: 170,
    align: 'center',
    format: (value: number) => value.toLocaleString('en-US'),
  },
];

interface Data {
  matchId: number;
  match: string;
  round: number;
  underway: boolean;
}

function createData(
  matchId: number,
  match: string,
  round: number,
  underway: boolean,
): Data {
  return { matchId, match, round, underway };
}

interface TableProps {
  matchData: Matches;
}

export default function StickyHeadTable({ matchData }: TableProps) {
  const rows: Data[] = matchData.match_list.map((match) => createData(match.id, `${match.player1_name} vs ${match.player2_name}`, match.round, match.underway));
 

  return (
    <Paper sx={{ width: '90%', overflow: 'hidden', borderColor: 'black', borderStyle: 'solid' }}>
      <TableContainer sx={{ maxHeight: 378 }}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              <TableCell colSpan={3} align='center'>{matchData.game_name}</TableCell>
            </TableRow>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows
              .map((row) => {
                return (
                  <TableRow hover role="checkbox" tabIndex={-1} key={row.matchId}>
                    <TableCell align='center' key='match'>
                      {row.match}
                    </TableCell>
                    <TableCell align='center' key='round'>
                      {row.round}
                    </TableCell>
                    <TableCell align='center' key='underway'
                       sx={{
                        color: row.underway ? 'green' : 'red',
                        fontSize: '1.2vw'
                      }}
                    >
                      {row.underway ? `Yes` : `No`}
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