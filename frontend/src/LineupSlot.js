import React from 'react';
import { useDrop } from 'react-dnd';
import { ItemTypes } from './ItemTypes';
import { IoPersonRemove } from 'react-icons/io5';

const LineupSlot = ({ index, player, movePlayerToSlot, removePlayerFromSlot }) => {
  const [{ isOver }, drop] = useDrop({
    accept: ItemTypes.PLAYER,
    drop: (item) => movePlayerToSlot(item, index),
    collect: (monitor) => ({
      isOver: !!monitor.isOver(),
    }),
  });

  return (
    <div ref={drop} className={`lineup-slot ${isOver ? 'is-over' : ''}`}>
      {index + 1}. {player ? <strong>{player.name}</strong> : ''} {player ? `(AVG: ${player.avg.toFixed(3)}, OBP: ${player.obp.toFixed(3)}, OPS: ${player.ops.toFixed(3)})` : ''}
      {player && <IoPersonRemove onClick={() => removePlayerFromSlot(index)} style={{ cursor: 'pointer', marginLeft: '10px' }} />}
    </div>
  );
};

export default LineupSlot;