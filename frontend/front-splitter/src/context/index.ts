'use client';

import { createContext } from 'react';
import { MemberList } from '../types';

export const MemberContext = createContext<MemberList>([]);
