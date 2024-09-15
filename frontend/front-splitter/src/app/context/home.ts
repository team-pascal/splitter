'use client';
import { createContext } from 'react';
import { MemberList } from '../types/type';

export const MemberContext = createContext<MemberList>([]);
