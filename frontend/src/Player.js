import React from 'react';
import { useDrag } from 'react-dnd';
import { ItemTypes } from './ItemTypes';

const Player = ({ player }) => {
  const [{ isDragging }, drag] = useDrag({
    type: ItemTypes.PLAYER,
    item: player,
    collect: (monitor) => ({
      isDragging: !!monitor.isDragging(),
    }),
  });

  return (
    <div ref={drag} className={`player ${isDragging ? 'is-dragging' : ''}`}>
      <strong>{player.name}</strong> (AVG: {player.avg.toFixed(3)}, OBP: {player.obp.toFixed(3)}, OPS: {player.ops.toFixed(3)})
    </div>
  );
};

export default Player;