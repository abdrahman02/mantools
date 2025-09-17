"use client";

import {
  ChevronRight,
  Database,
  FileText,
  Gauge,
  Home,
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
  SidebarGroupLabel,
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
import { useAuth } from "@/hooks/use-auth";

interface subNavMain {
  title: string;
  url: string;
  isActive?: boolean;
}

interface Nav {
  title: string;
  url: string;
  icon: React.ForwardRefExoticComponent<
    Omit<LucideProps, "ref"> & React.RefAttributes<SVGSVGElement>
  >;
  items?: subNavMain[];
}

const navAuth: Nav[] = [
  {
    title: "Summary",
    url: "/dashboard",
    icon: Gauge,
  },
];

const navMain: Nav[] = [
  { title: "Home", url: "/", icon: Home },
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
      {
        title: "PDF Tools",
        url: "/pdf-tools",
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
  {
    title: "Developer Tools",
    url: "#",
    icon: Database,
    items: [
      {
        title: "API Request Tester",
        url: "/api-request-tester",
      },
      {
        title: "JWT Decoder",
        url: "/jwt-decoder",
      },
      {
        title: "Hash Generator",
        url: "/hash-generator",
      },
    ],
  },
];

export function AppSidebar({
  accessToken,
  ...props
}: React.ComponentProps<typeof Sidebar> & { accessToken?: string }) {
  const pathname = usePathname();
  const { access, setAccess } = useAuth();

  // eslint-disable-next-line react-hooks/exhaustive-deps
  React.useEffect(() => setAccess(accessToken ?? null), [accessToken]);

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
        {navAuth.length > 0 && access && (
          <SidebarGroup>
            <SidebarGroupLabel>Dashboard</SidebarGroupLabel>
            <SidebarMenu>
              {navAuth.map((nav) => (
                <SidebarMenuItem key={nav.title}>
                  <SidebarMenuButton
                    className={
                      nav.url === pathname
                        ? "bg-main outline-border text-main-foreground"
                        : ""
                    }
                    tooltip={nav.title}
                    isActive={nav.url === pathname}
                    asChild
                  >
                    <Link href={nav.url}>
                      {nav.icon && <nav.icon />}
                      <span>{nav.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroup>
        )}
        <SidebarGroup>
          <SidebarMenu>
            {navMain.map((item) => {
              if (item.items === undefined) {
                const isActive = item.url === pathname;
                const className = isActive
                  ? "bg-main outline-border text-main-foreground"
                  : "";
                return (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      className={className}
                      tooltip={item.title}
                      isActive={isActive}
                      asChild
                    >
                      <Link href={item.url}>
                        {item.icon && <item.icon />}
                        <span>{item.title}</span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              }
              return (
                <Collapsible
                  key={item.title}
                  asChild
                  // defaultOpen={item.items?.some((itm) => itm.url === pathname)}
                  defaultOpen={true}
                  className="group/collapsible"
                >
                  <SidebarMenuItem>
                    <CollapsibleTrigger asChild>
                      <SidebarMenuButton
                        className={`data-[state=open]:bg-main data-[state=open]:outline-border data-[state=open]:text-main-foreground`}
                        tooltip={item.title}
                        isActive={item.items?.some(
                          (itm) => itm.url === pathname,
                        )}
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
              );
            })}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  );
}
