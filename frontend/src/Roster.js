import React from 'react';
import Player from './Player';

const Roster = ({ players }) => {
  return (
    <div>
      <h2>Roster</h2>
      {players.map((player, index) => (
        <Player key={index} player={player} />
      ))}
    </div>
  );
};

export default Roster;