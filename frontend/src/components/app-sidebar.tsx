"use client";

import {
  ChevronRight,
  FileText,
  Image,
  LucideProps,
  ThumbsUp,
} from "lucide-react";

import * as React from "react";

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarRail,
} from "@/components/ui/sidebar";
import Link from "next/link";
import { usePathname } from "next/navigation";

interface subNavMain {
  title: string;
  url: string;
  isActive?: boolean;
}

interface navMain {
  title: string;
  url: string;
  icon: React.ForwardRefExoticComponent<
    Omit<LucideProps, "ref"> & React.RefAttributes<SVGSVGElement>
  >;
  items: subNavMain[];
}

const navMain: navMain[] = [
  {
    title: "Text & Document",
    url: "#",
    icon: FileText,
    items: [
      {
        title: "Text Formatter",
        url: "/text-formatter",
      },
      {
        title: "Text Case Converter",
        url: "/text-case-converter",
      },
    ],
  },
  {
    title: "Image & Media",
    url: "#",
    icon: Image,
    items: [
      {
        title: "Images Compressor",
        url: "/images-compressor",
      },
      {
        title: "Images Converter",
        url: "/images-converter",
      },
      {
        title: "QR Generator",
        url: "/qr-generator",
      },
    ],
  },
];

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const pathname = usePathname();

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <Link href="/" className="group/header">
                <div className="bg-main text-main-foreground flex aspect-square size-8 items-center justify-center rounded-lg group-hover/header:border-2 group-hover/header:border-border">
                  <ThumbsUp className="size-4" />
                </div>
                <div className="grid flex-1 text-left leading-tight">
                  <h1 className="truncate text-sm">Mantools</h1>
                  <span className="truncate text-xs">Free online toolbox</span>
                </div>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarMenu>
            {navMain.map((item) => (
              <Collapsible
                key={item.title}
                asChild
                defaultOpen={item.items?.some((itm) => itm.url === pathname)}
                className="group/collapsible"
              >
                <SidebarMenuItem>
                  <CollapsibleTrigger asChild>
                    <SidebarMenuButton
                      className={`data-[state=open]:bg-main data-[state=open]:outline-border data-[state=open]:text-main-foreground`}
                      tooltip={item.title}
                      isActive={item.items?.some((itm) => itm.url === pathname)}
                    >
                      {item.icon && <item.icon />}
                      <span>{item.title}</span>
                      <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
                    </SidebarMenuButton>
                  </CollapsibleTrigger>
                  <CollapsibleContent>
                    <SidebarMenuSub>
                      {item.items?.map((subItem) => (
                        <SidebarMenuSubItem key={subItem.title}>
                          <SidebarMenuSubButton
                            isActive={subItem.url === pathname}
                            asChild
                          >
                            <Link href={subItem.url}>
                              <span>{subItem.title}</span>
                            </Link>
                          </SidebarMenuSubButton>
                        </SidebarMenuSubItem>
                      ))}
                    </SidebarMenuSub>
                  </CollapsibleContent>
                </SidebarMenuItem>
              </Collapsible>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  );
}
