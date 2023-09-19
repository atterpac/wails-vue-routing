import { Ref, ref } from "vue"
import { Column } from "../column"
import { EnumStringMember } from "@babel/types"

export type Fixture = {
    id: string,
    type?: string,
    number?: number,
    channel?: string,
    purpose?: string,
    position?: string,
    unit_number?: string,
    color?: string,
    dimmer?: string,
    ckt_num?: number,
    ckt_name?: string,
    system?: string,
    weight?: string,
    focus?: string,
    notes?: string
}

export const FixtureColumns: Ref<Column[]> = ref([
  {
      width: 100,
      data: "id",
      label: "ID",
      visible: false
  },{
      width: 30,
      data: "number",
      label: "Number",
      visible: false
  },{
      width: 30,
      data: "channel",
      label: "CH",
      visible: true
  },{
    width: 30,
    data: "type",
    label: "Type",
    visible: true
},{
      width: 30,
      data: "purpose",
      label: "Purpose",
      visible: true
  },{
      width: 100,
      data: "position",
      label: "Position",
      visible: true
  },{
      width: 100,
      data: "unit_number",
      label: "Unit",
      visible: true
  },{
      width: 100,
      data: "color",
      label: "Color",
      visible: true
  },{
      width: 100,
      data: "ckt_num",
      label: "Circuit",
      visible: true
  },{
      width: 100,
      data: "ckt_name",
      label: "Circuit Label",
      visible: true
  },{
      width: 100,
      data: "weight",
      label: "Weight",
      visible: true
  },{
      width: 100,
      data: "system",
      label: "System",
      visible: false
  },{
      width: 100,
      data: "focus",
      label: "Focus",
      visible: true
  }
  ])


export const FixtureData: Fixture[] = [
    {
      id: "1",
      type: "Moving Light",
      number: 1,
      channel: "1",
      purpose: "Front of House",
      position: "Left",
      unit_number: "12345",
      color: "White",
      dimmer: "100%",
      ckt_num: 1,
      ckt_name: "FOH",
      system: "ETC Element",
      weight: "100 lbs",
      focus: "Spot",
      notes: "None",
    },
    {
      id: "2",
      type: "Moving Light",
      number: 2,
      channel: "2",
      purpose: "Front of House",
      position: "Right",
      unit_number: "54321",
      color: "White",
      dimmer: "100%",
      ckt_num: 1,
      ckt_name: "FOH",
      system: "ETC Element",
      weight: "100 lbs",
      focus: "Spot",
      notes: "None",
    },
    {
      id: "3",
      type: "Wash Light",
      number: 3,
      channel: "3",
      purpose: "Front of House",
      position: "Center",
      unit_number: "67890",
      color: "White",
      dimmer: "100%",
      ckt_num: 1,
      ckt_name: "FOH",
      system: "ETC Element",
      weight: "50 lbs",
      focus: "Flood",
      notes: "None",
    },
    {
      id: "4",
      type: "Wash Light",
      number: 4,
      channel: "4",
      purpose: "Front of House",
      position: "Left",
      unit_number: "09876",
      color: "White",
      dimmer: "100%",
      ckt_num: 1,
      ckt_name: "FOH",
      system: "ETC Element",
      weight: "50 lbs",
      focus: "Flood",
      notes: "None",
    },
    {
      id: "5",
      type: "Wash Light",
      number: 5,
      channel: "5",
      purpose: "Front of House",
      position: "Right",
      unit_number: "98760",
      color: "White",
      dimmer: "100%",
      ckt_num: 1,
      ckt_name: "FOH",
      system: "ETC Element",
      weight: "50 lbs",
      focus: "Flood",
      notes: "None",
    },
    {
      id: "6",
      type: "Moving Light",
      number: 6,
      channel: "6",
      purpose: "Backlight",
      position: "Left",
      unit_number: "123456",
      color: "White",
      dimmer: "100%",
      ckt_num: 2,
      ckt_name: "Backlight",
      system: "ETC Element",
      weight: "100 lbs",
      focus: "Spot",
      notes: "None",
    },
    {
      id: "6",
      type: "Moving Light",
      number: 6,
      channel: "6",
      purpose: "Backlight",
      position: "Left",
      unit_number: "123456",
      color: "White",
      dimmer: "100%",
      ckt_num: 2,
      ckt_name: "Backlight",
      system: "ETC Element",
      weight: "100 lbs",
      focus: "Spot",
      notes: "None",
    },    {
        id: "1",
        type: "Moving Light",
        number: 1,
        channel: "1",
        purpose: "Front of House",
        position: "Left",
        unit_number: "12345",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "2",
        type: "Moving Light",
        number: 2,
        channel: "2",
        purpose: "Front of House",
        position: "Right",
        unit_number: "54321",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "3",
        type: "Wash Light",
        number: 3,
        channel: "3",
        purpose: "Front of House",
        position: "Center",
        unit_number: "67890",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "4",
        type: "Wash Light",
        number: 4,
        channel: "4",
        purpose: "Front of House",
        position: "Left",
        unit_number: "09876",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "5",
        type: "Wash Light",
        number: 5,
        channel: "5",
        purpose: "Front of House",
        position: "Right",
        unit_number: "98760",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },    {
        id: "1",
        type: "Moving Light",
        number: 1,
        channel: "1",
        purpose: "Front of House",
        position: "Left",
        unit_number: "12345",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "2",
        type: "Moving Light",
        number: 2,
        channel: "2",
        purpose: "Front of House",
        position: "Right",
        unit_number: "54321",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "3",
        type: "Wash Light",
        number: 3,
        channel: "3",
        purpose: "Front of House",
        position: "Center",
        unit_number: "67890",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "4",
        type: "Wash Light",
        number: 4,
        channel: "4",
        purpose: "Front of House",
        position: "Left",
        unit_number: "09876",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "5",
        type: "Wash Light",
        number: 5,
        channel: "5",
        purpose: "Front of House",
        position: "Right",
        unit_number: "98760",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },    {
        id: "1",
        type: "Moving Light",
        number: 1,
        channel: "1",
        purpose: "Front of House",
        position: "Left",
        unit_number: "12345",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "2",
        type: "Moving Light",
        number: 2,
        channel: "2",
        purpose: "Front of House",
        position: "Right",
        unit_number: "54321",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "3",
        type: "Wash Light",
        number: 3,
        channel: "3",
        purpose: "Front of House",
        position: "Center",
        unit_number: "67890",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "4",
        type: "Wash Light",
        number: 4,
        channel: "4",
        purpose: "Front of House",
        position: "Left",
        unit_number: "09876",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "5",
        type: "Wash Light",
        number: 5,
        channel: "5",
        purpose: "Front of House",
        position: "Right",
        unit_number: "98760",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },    {
        id: "1",
        type: "Moving Light",
        number: 1,
        channel: "1",
        purpose: "Front of House",
        position: "Left",
        unit_number: "12345",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "2",
        type: "Moving Light",
        number: 2,
        channel: "2",
        purpose: "Front of House",
        position: "Right",
        unit_number: "54321",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "3",
        type: "Wash Light",
        number: 3,
        channel: "3",
        purpose: "Front of House",
        position: "Center",
        unit_number: "67890",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "4",
        type: "Wash Light",
        number: 4,
        channel: "4",
        purpose: "Front of House",
        position: "Left",
        unit_number: "09876",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "5",
        type: "Wash Light",
        number: 5,
        channel: "5",
        purpose: "Front of House",
        position: "Right",
        unit_number: "98760",
        color: "White",
        dimmer: "100%",
        ckt_num: 1,
        ckt_name: "FOH",
        system: "ETC Element",
        weight: "50 lbs",
        focus: "Flood",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      },
      {
        id: "6",
        type: "Moving Light",
        number: 6,
        channel: "6",
        purpose: "Backlight",
        position: "Left",
        unit_number: "123456",
        color: "White",
        dimmer: "100%",
        ckt_num: 2,
        ckt_name: "Backlight",
        system: "ETC Element",
        weight: "100 lbs",
        focus: "Spot",
        notes: "None",
      }
]